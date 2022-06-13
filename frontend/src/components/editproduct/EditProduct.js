import React, { useState } from 'react'

const EditProduct = ({seenseen,productId}) => {
    const [productDetails,setProductDetails]=useState({
        name:"",model:"",specification:"",price:""
      });
      function handle(e){
         const dataFromInputFields={...productDetails }
         dataFromInputFields[e.target.id]=e.target.value
         setProductDetails(dataFromInputFields)
         console.log(dataFromInputFields)  
      } 
    // console.log(productId)
      async function editing(productId){
        // e.preventDefault();
     
          fetch("http://localhost:8080/product/"+productId, {
          method: "PUT",
          mode: "cors",
          headers: {
            "content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
          body:JSON.stringify(productDetails)
       
        })
          .then((res) => console.log(res))
          .catch((e) => console.log(e));
          console.log(productDetails)
          
      }

      if(seenseen){
        return (
            <div>
            
              <form onSubmit={()=>editing(productId)}>
                <label>productname</label>
                <input onChange={(e)=>handle(e)} value={productDetails.name} id="name" placeholder='product name' ></input>
                <label>model</label>
                <input  onChange={(e)=>handle(e)} value={productDetails.model } id="model" placeholder='model' ></input>
                <label>specification</label>
                <input onChange={(e)=>handle(e)} value={productDetails.specification } id="specification" placeholder='specification' ></input>
                <label>price</label>
                <input onChange={(e)=>handle(e)}  value={productDetails.price} id="price" placeholder='price' ></input>
                <button >edit</button>
              </form>
          
            </div>
          )

      }
      else{
        return null
      }
      
}

export default EditProduct