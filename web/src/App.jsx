import React from 'react';
import Header from "./components/Header/Header";

import './App.css';

class App extends React.Component {
    render() {
        return (
            <div className="App">
                <Header title="Baldur's Generator" />
            </div>
        );
    }
}

export default App;
