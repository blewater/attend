import * as React from "react";
import Item from "./Item";
import { fetchPlans, AddMember } from "../Members/helpers";
import ItemDetails from "./Itemdetails";
import MaterialTableComponent from "../UI/MaterialTable";
import RadioGroupWithLabels from "../UI/RadioGroupWithLabels";

interface IAppProps {}

const List: React.FC<IAppProps> = props => {
  const [plans, setPlans] = React.useState([]);
  const [itemForm, setItemForm] = React.useState(false);
  const [formStatus, setFormStatus] = React.useState("add");
  const [selectedMember, setselectedMember] = React.useState({});

  // Similar to componentDidMount and componentDidUpdate:
  React.useEffect(() => {
   // fetchData();
    // Update the document title using the browser API
  }, []);

  const fetchData = () =>
  fetchPlans().then(response => {
      let columns = [];
      let data = [];
      Object.keys(response.data[0]).map((key, i) => {
        columns.push({ title: key.toUpperCase(), field: key });
      });
      response.data.map((d, i) => {
       data.push(d);
      });
      let setData = {
        columns: columns,
        data: data
      };
      setPlans(setData);
    });

  const handleAdditem = e => {
    e.preventDefault();
    setItemForm(true);
    setFormStatus("add");
   // AddMember();
  };

  const handleSelectedItem = (id: string) => {
    let selectedItem = Plans.filter(plan => {
      return plan.PlanID === id;
    });
    console.log("SelectedItem", selectedItem);
    setFormStatus("edit");
    setselectedMember(selectedItem[0]);
    setItemForm(true);
  };

  return (
    <div>
      {plans ? (
        <>
        <br/>
        {/* <RadioGroupWithLabels/> */}
        {/* <MaterialTableComponent Plans={plans} /> */}
        </>
      ) : (
        <div>Loading Data..</div>
      )}

      {/* {itemForm ? (
     <ItemDetails Plans={Plans[0]} member={selectedMember} status={formStatus} />
      ) : null} */}
    </div>
  );
};

export default List;
