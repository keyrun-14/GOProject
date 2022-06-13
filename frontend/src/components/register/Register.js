import React, { useState } from "react";
import "./register.css"
const Register = () => {
const [register,setRegister]=useState({
  name:"",email:"",password:""
  // ,confirmPassword:""
})
function handlingInput(e){
  const inputData = {...register}
   inputData[e.target.id] = e.target.value
   setRegister(inputData)
   console.log(inputData)
}
function registration_submission(e){
  e.preventDefault();

  fetch("http://localhost:8080/register",{
    method:"POST",
    mode:"cors",
    headers:{"content-Type":"application/json",'Access-Control-Allow-Origin':'*'},
    body:JSON.stringify(register)
  }).then((res)=>console.log(res)).catch(err=>console.log(err))
  window.location.href="http://localhost:3000/"

}

  return (
    <div className="registration">
      <div className="registraion-form-container">
      <form className="register-form" onSubmit={(e)=>registration_submission(e)} method="POST">
      <h1>Registration</h1>
        <div>
          <label>User Name : </label>
          <input className="register-input-fields" onChange={(e)=>handlingInput(e)} id="name" value={register.name} type="text" placeholder="enter User Name" required></input>
        </div>
        <div>
          <label>Email : </label>
          <input className="register-input-fields" onChange={(e)=>handlingInput(e)} id="email" value={register.email} type="email" placeholder="enter email" required></input>
        </div>
        <div>
          <label>Password : </label>
          <input className="register-input-fields" onChange={(e)=>handlingInput(e)} id="password" value={register.password} type="password" placeholder="enter Password" required></input>
        </div>
        {/* <div>
          <label>Confirm Password : </label>
          <input className="register-input-fields" onChange={(e)=>handlingInput(e)} id="confirmPassword" value={register.confirmPassword} type="password" placeholder="Confirm Password" required></input>
        </div> */}
        <div>
        <button>Register</button>
        </div>
          
        
      </form>

      </div>
      
    </div>
  );
};

export default Register;
