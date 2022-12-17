import * as ff from '@google-cloud/functions-framework'

ff.http('helloWorld', (req: ff.Request, res: ff.Response) => {
  res.send('Hello World')
})
