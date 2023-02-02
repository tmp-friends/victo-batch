import { hashtags, tweet_objects, media_objects } from '@prisma/client';
import {
  TwitterApi,
  TwitterApiReadOnly,
  TweetSearchRecentV2Paginator,
} from "twitter-api-v2"

import { DaoFanartTweet } from "./dao-fanart-tweet"

export class ServiceFanartTweet {
  private roClient: TwitterApiReadOnly
  private dao: DaoFanartTweet

  constructor() {
    const BEARER_TOKEN = process.env.BEARER_TOKEN ?? ""
    const client = new TwitterApi(BEARER_TOKEN)
    this.roClient = client.readOnly

    this.dao = new DaoFanartTweet()
  }

  /**
   * @remarks
   * hashtagsテーブルのカラムをすべて取得
   *
   * @returns hashtags[]
   */
  public async getHashtags(): Promise<hashtags[]> {
    return await this.dao.getHashtags()
  }

  /**
   * @remarks
   * TwitterAPIを用いて前日分のツイートを取得する処理
   *
   * @param hashtagName - 取得したいツイートに付与されているハッシュタグ名
   * @returns 取得したツイート配列
   */
  public async fetchTweets(
    hashtagName: string,
  ): Promise<TweetSearchRecentV2Paginator> {
    const searchKeyword = `#${hashtagName} -is:retweet has:images`
    const [yesterdayMidnight, todayMidnight] = await this.setWithinTime()

    const fanartTweets = await this.roClient.v2.search(searchKeyword, {
      start_time: yesterdayMidnight,
      end_time: todayMidnight,
      expansions: ['author_id', 'attachments.media_keys'],
      'tweet.fields': ['created_at', 'public_metrics'],
      'media.fields': ['preview_image_url', 'url', 'variants'],
    })

    return fanartTweets
  }

  /**
   * @remarks
   * (日本標準時間で)前日の開始と終了の時刻を取得する処理
   *
   * @returns [yesterdayMidnight, todayMidnight] - [前日深夜0時, 本日深夜0時]
   */
  private async setWithinTime(): Promise<string[]> {
    // 日本標準時間で取得
    const date = new Date(
      Date.now() + (new Date().getTimezoneOffset() + 9 * 60) * 60 * 1000,
    )

    // ISO規格
    // getMonth()は0はじまりのため1加算する必要あり
    const yesterdayMidnight =
      `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate() - 1}T00:00:00+09:00`
    const todayMidnight =
      `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}T00:00:00+09:00`

    return [yesterdayMidnight, todayMidnight]
  }

  /**
   * @remarks
   * 取得したツイートを整形しDBにinsertする処理
   *
   * @param id - hashtagsテーブルのID
   * @param fanartTweets - twitter-api-v2で取得したツイート配列
   * @returns void
   */
  public async insertTweets(
    id: number,
    fanartTweets: TweetSearchRecentV2Paginator,
  ): Promise<void> {
    const data: tweet_objects[] = []
    for await (const fanartTweet of fanartTweets) {
      const [text, tweetUrl] = await this.extractTweetUrl(fanartTweet['text'])

      const tweetData: tweet_objects = {
        tweet_id: fanartTweet['id'],
        text: text,
        retweet_count: Number(fanartTweet['public_metrics']['retweet_count']),
        like_count: Number(fanartTweet['public_metrics']['like_count']),
        author_id: fanartTweet['author_id'],
        tweet_url: tweetUrl,
        tweeted_at: new Date(fanartTweet['created_at']),
        created_at: new Date(),
        updated_at: new Date(),
        hashtag_id: id,
      }
      data.push(tweetData)
    }

    await this.dao.insertTweetObjects(data)
  }

  /**
   * @remarks
   * ツイート本文に含まれているツイートURLを抽出する処理
   *
   * @params tweetText - ツイート本文
   * @returns [text, tweetUrl] - [ツイート本文, ツイートURL]
   */
  private async extractTweetUrl(tweetText: string): Promise<string[]> {
    // https://t.co/<空白以外の1文字以上>
    const tweetUrl = tweetText.match(/https:\/\/t\.co\/\S*$/)[0]
    const text = tweetText.replace(` ${tweetUrl}`, '')

    return [text, tweetUrl]
  }

  /**
   * @remarks
   * 取得したツイートのメディアを整形しDBにinsertする処理
   *
   * @param fanartTweets - twitter-api-v2で取得したツイート配列
   * @returns void
   */
  public async insertMedia(
    fanartTweets: TweetSearchRecentV2Paginator
  ): Promise<void> {
    const data: media_objects[] = []
    for await (const fanartTweet of fanartTweets) {
      // media.fieldsは追加情報
      const mediaFields = fanartTweets.includes.medias(fanartTweet)
      // 画像は複数枚ある
      for await (const mediaField of mediaFields) {
        const mediaData: media_objects = {
          type: mediaField['type'],
          // 画像以外のmedia_typeも考慮
          url: mediaField['url'] ?? mediaField['variants'][0]['url'],
          created_at: new Date(),
          updated_at: new Date(),
          tweet_id: fanartTweet['id'],
        }
        data.push(mediaData)
      }
    }

    await this.dao.insertMediaObjects(data)
  }
}
