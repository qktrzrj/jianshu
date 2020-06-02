import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import {ApolloProvider} from '@apollo/react-hooks';
import {createUploadLink} from 'apollo-upload-client'
import {ApolloClient} from "apollo-client";
import {InMemoryCache} from "apollo-cache-inmemory";

const client = new ApolloClient({
    link: createUploadLink({
        uri: 'http://localhost:8008/graphql',
        credentials: 'include',
    }),
    cache: new InMemoryCache(),
});

ReactDOM.render(
    <ApolloProvider client={client}>
        <App/>
    </ApolloProvider>,
    document.getElementById('root')
);