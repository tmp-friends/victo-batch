import * as ff from '@google-cloud/functions-framework'

import { LogicFanartTweet } from './insert-fanart-tweet/logic-fanart-tweet'

// Cloud Functionsのエントリーポイント
ff.http('insertFanartTweets', async (req: ff.Request, res: ff.Response) => {
  const logic = new LogicFanartTweet()
  await logic.doExecute()

  // processが強制終了しなければ正常終了
  res.status(200).send("Success");
})
