window.onload = () => {
  var pattern =
    /^[a-zA-Z0-9]+([-_.][A-Za-zd]+)*@([a-zA-Z0-9]+[-.])+[A-Za-zd]{2,5}$/;
  var mail = document.getElementsByName("ml")[0];
  var password = document.getElementsByName("pw")[0];
  var apassword = document.getElementsByName("apw")[0];
  var code = document.getElementsByName("cd")[0];
  // console.log(password);
  var reg = document.getElementById("sb");

  //   var tp = document.getElementsByTagName("p")[0];
  // console.log(tp.innerHTML)

  //   mail.oninput = () => {
  //     mail.addEventListener("keydown", () => {});
  //   };

  reg.addEventListener("click", function a() {
    if (!mail.value.match(pattern)) {
      // console.log(mail.value);
      alert("邮箱有误，请重试");
      //   window.location.href = "#"
      window.navigator("#");
    }
  });
};
