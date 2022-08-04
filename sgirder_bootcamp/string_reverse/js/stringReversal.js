const strToReverse = "hello";



function revStr(s) {
    var t = "";
    for (let i = 0; i < s.length; i++){
        t = s[i] + t
    };
    return t

};

console.log(revStr(strToReverse));

exports.revStr = revStr;