export interface IData {
  data:IState
}

export interface IState {
  characters: ICharacter[];
  info: Iinfo[];
  loading: boolean;
  species: string[];
  errorInSearch:boolean;
}

export interface Iinfo {
  count: number;
  pages: number;
  next: string;
  prev: string;
}
export interface IMember {
  MembersID: number;
  SurName: string;
  SLEN: number;
  Name: string;
  Father?: any;
  Job?: any;
  Age: number;
  Phone1: string;
  EMail?: any;
  Area: string;
  IdentityCard: string;
  Photo?: any;
  Status: number;
  StatusD: string;
  CloseD: Date;
  SDR?: any;
  SCR?: any;
  BAL?: any;
  LastVisit?: any;
}
