// get token
import { appUrl } from "../config";

let user = {
  username: "gmlogic@gmail.com",
  password: "1Mgergm++",
};

let token: string;

export const getToken = async (creds:any) => {
  let response = await fetch("http://api.rem4you.com/api/Token", {
    method: "POST",
    headers: {
      "Content-Type": "application/json;charset=utf-8",
    },
    body: JSON.stringify(creds),
  });
  let result = await response.json();
  console.log(result);
  token = result.access_token;
  if(token) {
    localStorage.setItem("AppToken", token);
    return true
  }
  return false;
};
//curl -X GET --header 'Accept: application/json' --header 'Authorization: bearer eyJhbGciOiJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGRzaWctbW9yZSNobWFjLXNoYTI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1laWRlbnRpZmllciI6IjZlYWRiZDAyLWQ5ODktNGEyYy1hY2NjLTRjNGY2YmU0ZDIxNiIsImh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3dzLzIwMDUvMDUvaWRlbnRpdHkvY2xhaW1zL25hbWUiOiJnbWxvZ2ljQGdtYWlsLmNvbSIsImh0dHA6Ly9zY2hlbWFzLm1pY3Jvc29mdC5jb20vYWNjZXNzY29udHJvbHNlcnZpY2UvMjAxMC8wNy9jbGFpbXMvaWRlbnRpdHlwcm92aWRlciI6IkFTUC5ORVQgSWRlbnRpdHkiLCJBc3BOZXQuSWRlbnRpdHkuU2VjdXJpdHlTdGFtcCI6Ijg2MzYzODA5LTQyYTQtNDdlMS1hMGIyLTgyOWUxYzFmZDNhMCIsImh0dHA6Ly9zY2hlbWFzLm1pY3Jvc29mdC5jb20vd3MvMjAwOC8wNi9pZGVudGl0eS9jbGFpbXMvcm9sZSI6WyJTdXBlckFkbWluIiwiVXNlciIsIkFkbWluIl0sIlBob25lIjoiNjk3MzM3NTMwMyIsIkdlbmRlciI6Ik1hbGUiLCJuYmYiOjE1ODY0NjgzMzIsImV4cCI6MTU4NjcyNzUyOCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1NDg5OCIsImF1ZCI6IjA5OTE1M2MyNjI1MTQ5YmM4ZWNiM2U4NWUwM2YwMDIyIn0.rbDCi6u61hhv95axbv4Mn28LARkvcralVTS8Df8i458' 'http://api.rem4you.com/members/obj?pars=%20%7BSearchValue%3A%20%22%CE%A3%CE%A4%CE%95%CE%9B%22%2CBOption%3A%200%2CDFrom%3A%20%2212%2F31%2F2013%22%2CDTo%3A%20%2212%2F31%2F2013%22%2CTakeRecs%3A0%2CId%3A0%2CLastId%3A0%7D'
//"http://api.rem4you.com/members/obj?pars=%7B%22SearchValue%22%3A%22%CE%A3%CE%A4%CE%95%CE%9B%22%2C%22BOption%22%3A0%2C%22DFrom%22%3A%2212%2F31%2F2013%22%2C%22DTo%22%3A%2212%2F31%2F2013%22%2C%22TakeRecs%22%3A0%2C%22Id%22%3A0%2C%22LastId%22%3A0%7D"

//"http://api.rem4you.com/members/obj?pars=%7B%22SearchValue%22%3A%22%CE%A3%CE%A4%CE%95%CE%9B%22%2C%22BOption%22%3A0%2C%22DFrom%22%3A%2212%2F31%2F2013%22%2C%22DTo%22%3A%2212%2F31%2F2013%22%2C%22TakeRecs%22%3A0%2C%22Id%22%3A0%2C%22LastId%22%3A0%7D"
export const makeRequest = async (
  method: string = "POST",
  url: string = "https://api.rem4you.com/api/Token",
  body: any = user
) => {
  const auth = localStorage.getItem("AppToken");
  if (auth) {
    const defaultHeaders = {
      Accept: 'application/json',
      Authorization: "bearer " + auth,
    };
    const response = await fetch(url, {
      method: method,
      headers: defaultHeaders,
    });
    let result = await response.json();
    return result;
  }
};

