import * as React from "react";

interface IAppProps {
  data: any;
  editItem:Function;
}

const Item: React.FunctionComponent<IAppProps> = props => {
  const data = { ...props.data };


  const handleMember=(e)=>{
      e.preventDefault()
    props.editItem(data.MembersID)
  }


  return (
    
    <tr>
      <td>
          <div className="d-flex">
        <button onClick={handleMember}>Edit</button>
        <button>Delete</button></div>
      </td>
      <td>
        <span>{data.SurName}</span>
      </td>
      <td>
        <span>{data.Name}</span>
      </td>
      <td>
        <span>{data.Phone1}</span>
      </td>
      <td>
        <span>{data.Emai}dfsdfsdf</span>
      </td>
      <td>
        <span>{data.Status}</span>
      </td>
    </tr>
  );
};

export default Item;
