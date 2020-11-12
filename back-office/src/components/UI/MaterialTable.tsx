import * as React from 'react';
import MaterialTable, { Column } from 'material-table';

import { AddMember } from "../Members/helpers"


interface TableState {
  columns: Array<Column<Row>>;
  data: Row[];
}

export default function MaterialTableComponent(props) {
  const [state, setState] = React.useState<TableState>({});


    // Similar to componentDidMount and componentDidUpdate:
    React.useEffect(() => {
        setState(props.members);
        // Update the document title using the browser API
      }, [props.members]);

  console.log(props)

  return (
    <MaterialTable
      title="Members"
      columns={state.columns}
      data={state.data}
      options={{
        pageSize:10
      }}
      editable={{
        onRowAdd: newData =>
          new Promise(resolve => {
            setTimeout(() => {
              resolve();
              setState(prevState => {
                const data = [...prevState.data];
                console.log(newData)
                AddMember(newData)
                    data.push(newData);
                    return { ...prevState, data };
                
              });
            }, 600);
          }),
        onRowUpdate: (newData, oldData) =>
          new Promise(resolve => {
            setTimeout(() => {
              resolve();
              if (oldData) {
                setState(prevState => {
                  const data = [...prevState.data];
                  data[data.indexOf(oldData)] = newData;
                  return { ...prevState, data };
                });
              }
            }, 600);
          }),
        onRowDelete: oldData =>
          new Promise(resolve => {
            setTimeout(() => {
              resolve();
              setState(prevState => {
                const data = [...prevState.data];
                data.splice(data.indexOf(oldData), 1);
                return { ...prevState, data };
              });
            }, 600);
          }),
      }}
    />
  );
}