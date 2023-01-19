import { hashtags, media_objects, PrismaClient, tweet_objects } from '@prisma/client'

export class DaoFanartTweet {
  private prisma: PrismaClient

  constructor() {
    this.prisma = new PrismaClient()
  }

  /**
   * @remarks
   * hashtagsテーブルのカラムをすべて取得
   *
   * @returns hashtags[]
   */
  public async getHashtags(): Promise<hashtags[]> {
    return await this.prisma.hashtags.findMany()
  }

  /**
   * @remarks
   * tweet_objects[]をDBにinsertする処理
   *
   * @param tweet_objects[] - twitter-api-v2で取得したツイート配列
   * @returns void
   */
  public async insertTweetObjects(data: tweet_objects[]): Promise<void> {
    await this.prisma.tweet_objects.createMany({ data, skipDuplicates: true })
  }

  /**
   * @remarks
   * media_objects[]をDBにinsertする処理
   *
   * @param media_objects[] - twitter-api-v2で取得したツイートのメディア配列
   * @returns void
   */
  public async insertMediaObjects(data: media_objects[]): Promise<void> {
    await this.prisma.media_objects.createMany({ data, skipDuplicates: true })
  }
}
