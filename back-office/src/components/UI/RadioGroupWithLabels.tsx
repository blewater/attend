import * as React from "react";
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';

export default function RadioGroupWithLabels(props) {
  const [value, setValue] = React.useState('female');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValue((event.target as HTMLInputElement).value);
  };

  return (
    <FormControl component="fieldset">
      <RadioGroup aria-label="position" name="position" value={value} onChange={handleChange} row>
        <FormControlLabel
          value="top"
          control={<Radio color="primary" />}
          label="All"
          labelPlacement="top"
        />
        <FormControlLabel
          value="start"
          control={<Radio color="primary" />}
          label="Active"
          labelPlacement="top"
        />
        <FormControlLabel
          value="bottom"
          control={<Radio color="primary" />}
          label="Inactive"
          labelPlacement="top"
        />
      </RadioGroup>
    </FormControl>
  );
}
