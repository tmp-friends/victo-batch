import { ILogObj } from 'tslog/dist/types/interfaces';
import { Logger } from 'tslog'
import * as dotenv from "dotenv"

import { ServiceFanartTweet } from "./service-fanart-tweet"

export class LogicFanartTweet {
  private logger: Logger<ILogObj>
  private service: ServiceFanartTweet

  constructor() {
    this.logger = new Logger({ name: "insertFanartTweets" })

    // 環境変数の読み込み
    dotenv.config()

    this.service = new ServiceFanartTweet()
  }

  /**
   * @remarks
   * ツイートを取得し、DBにインサートするバッチ処理
   */
  public async doExecute() {
    this.logger.info("Start")

    try {
      const hashtagList = await this.service.getHashtags()

      for await (const hashtag of hashtagList) {
        const fanartTweets = await this.service.fetchTweets(hashtag['tagName'])
        this.logger.debug(`${hashtag["tagName"]}: ${fanartTweets.tweets.length}`)

        await this.service.insertTweets(hashtag['id'], fanartTweets)
        await this.service.insertMedia(fanartTweets)
      }
      this.logger.info("Success")
    } catch (e) {
      this.logger.fatal(e)
      this.logger.info("Failed")
      // 強制終了
      process.exit(1)
    }
  }
}
