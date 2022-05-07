const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

module.exports.getCommand = function (nc, subject, sc) {
    //let videoPath = "/Users/alifarhadi/Desktop/Codes/Interview/Uniqcast/task/video.mp4"
    rl.question('Enter video path: ', function (videoPath) {
        nc.publish(subject, sc.encode(videoPath));
    });
    rl.on('error', function (e) {
        nc.drain();
        return e
    });
};