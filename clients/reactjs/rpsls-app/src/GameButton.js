import React from 'react';
import './GameButton.css';

var lastSelectedId = 0;
//https://www.tutorialspoint.com/reactjs/reactjs_environment_setup.htm
//localhost:3000
//https://github.com/facebook/create-react-app
export class ChoiceButton extends React.Component {
    constructor(props) {
      super(props);
      this.state = { picked: false, name: "none"};
    }
  
    render() {
      return React.createElement(
        'button'
        ,{ onClick: () => {this.setState( { picked: true } ); lastSelectedId = this.props.id; console.log("last selected: ", lastSelectedId)} }
        ,this.props.name
      );
    }
  }
  
  // You must select a decision...
  // Reference - https://selectadecision.info/woods.html
  export class SelectADecision extends React.Component {
    constructor(props) {
      super(props);
      this.state = { picked: false, menu: []};
    }
  
    //https://blog.hellojs.org/fetching-api-data-with-react-js-460fe8bbf8f2
    componentDidMount() {
      fetch('https://localhost:4077/choices')
      .then( results => {
        return results.json();
      }).then( data => {
        console.log("Results: ", data)
        let options = data.map((option) => {
            return <ChoiceButton name={option.name} id={option.id} key={option.id} />
        })
          this.setState({menu : options});
          console.log("state", this.state.menu)
        })
    }
    render () {
      console.log(this)
      return (this.state.menu);
    }
  }

  export default ChoiceButton 