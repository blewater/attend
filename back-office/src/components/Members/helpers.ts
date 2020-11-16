import 'babel-polyfill'
import {makeRequest} from '../../api/fetch'
import { appUrl } from "../../config";

let params1:any = {SearchValue: "ΣΤΕΛ",BOption: 0,DFrom: "12/31/2013",DTo: "12/31/2013",TakeRecs:0,Id:0,LastId:0}

params1=JSON.stringify(params1)
export const  fetchMembers= async ()=>{

  return await makeRequest("GET",`http://api.rem4you.com/members/obj?pars=${encodeURIComponent(params1)}`).then(response=>{
      console.log("MEMBERS",response)
      return response
    })
    
}
