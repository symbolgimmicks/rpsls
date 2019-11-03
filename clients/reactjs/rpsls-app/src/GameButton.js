import React from 'react';
import './GameButton.css';

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
        'button',
        { onClick: () => this.setState( { picked: true } ) }
        ,this.state.name
      );
    }
  }
  
  export class rockButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "rock"}
    }
  }
  
  export class paperButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "paper"}
    }
  }
  
  export class scissorsButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "scissors"}
    }
  }
  
  export class lizardButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "lizard"}
    }
  }
  
  export class spockButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "spock"}
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
          return (
            // format your returned data into a map to use for rendering later.
            
            //React.createElement("div",{value: option.id, id:option.name+"_button"},)
                (<div id= {option.name}>{option.name}</div>)
            );
        })
          this.setState({menu : options});
          console.log("state", this.state.menu)
        })
    }
    render () {
      console.log(this)
      return (
        this.state.menu
        // tODO This is where you might make the menu items.
      );
    }
  }

  export default ChoiceButton 