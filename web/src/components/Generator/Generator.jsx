import React from 'react';

class Generator extends React.Component {
    constructor(props) {
        super(props);
        this.state = { name: null };

        this.charData = React.createRef();
    }

    generateData = () => {
        // TODO: Fetch data from baldur's gen api
        this.setState({
            'name': 'Conner'
        });
    }

    render() {
        return (
            <div className="Generator">
                <button onClick={this.generateData}>NEW CHARACTER</button>
                <div id="CharacterData" ref={this.charData}>
                    <p className="Name">{this.state.name}</p>
                </div>
            </div>
        );
    }
}

export default Generator;