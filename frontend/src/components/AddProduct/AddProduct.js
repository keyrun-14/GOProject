import React, {  useState } from 'react'

const AddProduct = () => {
  // const [productDetails,setProductDetails]=useState({
  //   name:"",model:"",specification:"",price:""
  // });
  const [name,setName]=useState("")
  const [model,setModel]=useState("")
  const [specification,setSpecification]=useState("")
  const [price,setPrice]=useState("")
  // function handle(e){
  //   //  const dataFromInputFields={...productDetails }
  //   //  dataFromInputFields[e.target.id]=e.target.value
  //    setProductDetails({...productDetails,[e.target.id]:e.target.value})
 
  // } 

  async function adding(e){
    // e.preventDefault();
 
      fetch("http://localhost:8080/product", {
      method: "POST",
      mode: "cors",
      headers: {
        "content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
      body:JSON.stringify({name:name,model:model,specification:specification,price:price})
   
    })
      .then((res) => res.json())
      .then((data)=>console.log(data))
     


   
  }
  return (
    <div>
      <form onSubmit={(e)=>adding(e)}>
        <table>
          <thead>
          <tr>
            <th>
               {/* <label>productname</label> */}
        <input onChange={(e)=>{setName(e.target.value)}} value={name} id="name" placeholder='product name' required></input>

            </th>
            <th>{/* <label>model</label> */}
        <input  onChange={(e)=>{{setModel(e.target.value)}}} value={model } id="model" placeholder='model' required></input> </th>
               
        <th>   {/* <label>specification</label> */}
        <input onChange={(e)=>{{setSpecification(e.target.value)}}} value={specification } id="specification" placeholder='specification' required></input></th>
       <th> {/* <label>price</label> */}
        <input onChange={(e)=>{{setPrice(e.target.value)}}}  value={price} id="price" placeholder='price' required></input>  </th>
       
       

          </tr>
          

          </thead>
          
        

        </table>
        <button onClick={adding}>add</button>
      
      </form>
    </div>
  )
}

export default AddProduct