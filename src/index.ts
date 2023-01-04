import * as ff from '@google-cloud/functions-framework'

import { Logic } from './fanart-tweet/logic'

ff.http('insertFanartTweets', (req: ff.Request, res: ff.Response) => {
  const logic = new Logic()
  logic.doExecute()
})
