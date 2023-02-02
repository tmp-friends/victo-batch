import * as ff from '@google-cloud/functions-framework'

import { LogicFanartTweet } from './insert-fanart-tweets/logic-fanart-tweet'

/**
 * Cloud Functionsのエントリーポイント
 * --targetオプションで実行したい関数を指定する
 */
ff.http('registerHashtags', async (req: ff.Request, res: ff.Response) => {
  const logic = new LogicFanartTweet()
  await logic.doExecute()

  // processが強制終了しなければ正常終了
  res.status(200).send("Success");
})

ff.http('insertFanartTweets', async (req: ff.Request, res: ff.Response) => {
  const logic = new LogicFanartTweet()
  await logic.doExecute()

  // processが強制終了しなければ正常終了
  res.status(200).send("Success");
})
