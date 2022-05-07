const nats = require('nats');
const sc = nats.StringCodec();
var cmd = require("./cmd");
var configs = require("./configs");

module.exports.provideNatsConnection = async function () {
  const nastServer = configs._Url;
  const requestSubject = configs._RequestSubject;
  const responseSubject = configs._ResponseSubject;

  const nc = await nats.connect(
    {
      servers: nastServer
    }
  );
  module.exports.nc;

  const sub = nc.subscribe(responseSubject);

  (async () => {
    for await (const m of sub) {
      console.log(`The output file has been saved to: ${sc.decode(m.data)}\n`);
      cmd.getCommand(nc, requestSubject, sc);
    }
  })();

  cmd.getCommand(nc, requestSubject, sc);
};