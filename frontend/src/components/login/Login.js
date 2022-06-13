import React, { useState } from "react";
import "./Login.css"
const Login = () => {
  const [login,setLogin]=useState({
    email:"",password:""
  })
  function handling(e){
    const login_data={...login}
    login_data[e.target.id]=e.target.value
    setLogin(login_data)
    console.log(login_data)
  }
async function login_submission(e){
  e.preventDefault();
  await fetch("http://localhost:8080/login",{
    method:"POST",
    mode:"cors",
    headers:{"content-Type":"application/json",'Access-Control-Allow-Origin':'*'},
    body:JSON.stringify(login)
  }).then(()=> localStorage.setItem("authtoken",)).catch(()=>window.location.href="http://localhost:3000/")
  
 

 }

  return (
    <div>
      <div className="Login">
        <div className="form-container">
          <form className="Login-form" onSubmit={(e)=>login_submission(e)}>
            <h1>Login</h1>
            <div>
              <label>Email : </label>
              <input className="login-input-fields" onChange={(e)=>handling(e)} id="email" value={login.email} type="email" placeholder="enter email" required></input>
            </div>
            <div>
              <label>Password : </label>
              <input className="login-input-fields" onChange={(e)=>handling(e)} id="password" value={login.password} type="password"  placeholder="enter Password"   required ></input>
            </div >
            <div className="button-div">
            <button>Login</button>
            </div>
            
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
