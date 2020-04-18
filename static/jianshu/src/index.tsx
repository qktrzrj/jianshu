import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import ApolloClient from 'apollo-boost';
import {ApolloProvider} from '@apollo/react-hooks';

const client = new ApolloClient({
    uri: 'http://localhost:8008/graphql',
    credentials:'include',
    onError: ({networkError}) => {
        // @ts-ignore
        if (networkError && networkError.statuscode === 401) {
            window.open("/signIn")
        }
    },
});

ReactDOM.render(
    <ApolloProvider client={client}>
        <App/>
    </ApolloProvider>,
    document.getElementById('root')
);