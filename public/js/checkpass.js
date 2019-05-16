function CheckPassword(){
	 var i = document.login.confirm.value
	 var j = document.login.password.value
	if(i != j){
		window.alert("password does not much")
		return false
	}else{
		return true
	}
}
