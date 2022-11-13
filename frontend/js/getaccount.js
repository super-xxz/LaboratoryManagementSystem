//测试一下效果 因为具体账号需要后端调数据库的

window.onload = ()=>{

    var userid = document.getElementsByTagName("p")[0];

//生成随机账号
    function getRandomCode(length) {
        if (length > 0) {
        var data = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"];
        var nums = "";
        for (var i = 0; i < length; i++) {
            var r = parseInt(Math.random() * 36);
            nums += data[r];
        }
        return nums;
        } 
        
        else {
        return false;
        }
    }
    
    var res = getRandomCode(6);
    // console.log(userid);

    userid.innerHTML = res;

}

