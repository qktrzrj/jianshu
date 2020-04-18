import {gql} from "apollo-boost";

export const CheckUsernameGQL = gql`
query ValidUsername($username:String!){
  ValidUsername(username:$username)
}
`;

export const CheckEmailGQL = gql`
query ValidEmail($email:String!){
  ValidEmail(email:$email)
}
`;

export const SignUpGQL  = gql`
mutation SignUp($email:String!,$password:String!,$username:String!){
  SignUp(email:$email,password:$password,username:$username){
    id
    username
    introduce
    avatar
    state
    root
  }
}
`;

export  const SignInGQL = gql`
mutation SignIn($username:String!,$password:String!,$rememberme:Boolean!){
  SignIn(username:$username,password:$password,rememberme:$rememberme){
    id
    username
    introduce
    avatar
    state
    root
  }
}
`