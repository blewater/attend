import * as React from "react";

interface IAppProps {
}

const ItemDetails: React.FunctionComponent<IAppProps> = props => {

    let item = props.status==="edit"? props.member:props.members
    console.log("Member",props.member)
  
  return (

   <form  className={'d-flex f-wrap m-1'}>
         {Object.keys(item).map((key,i)=>{
           return <div key={key+i}  className={'d-flex'} > 
             <label style={{ width: "80px" }} className={'mx-1'} htmlFor={key+i} key={key+i}>{key}</label>
           <input name={key} value={props.status==="edit"?props.member[key]:''}  id={key+i}/>
           </div>
         })}
         <div className={"m-1"} style={{width:'100%'}}><button>Save</button></div>
       
   </form>
  );
};

export default ItemDetails;
