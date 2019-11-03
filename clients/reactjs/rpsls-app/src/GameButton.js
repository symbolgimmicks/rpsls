import React from 'react';
import './GameButton.css';

//https://www.tutorialspoint.com/reactjs/reactjs_environment_setup.htm
//localhost:3000
//https://github.com/facebook/create-react-app
class ChoiceButton extends React.Component {
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
  
  class RockButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "rock"}
    }
  }
  
  class PaperButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "paper"}
    }
  }
  
  class ScissorsButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "scissors"}
    }
  }
  
  class LizardButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "lizard"}
    }
  }
  
  class SpockButton extends ChoiceButton {
    constructor(props) {
      super(props);
      this.state =  { name: "spock"}
    }
  }
  
  // You must select a decision...
  // Reference - https://selectadecision.info/woods.html
  class SelectADecision extends React.Component {
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
              
              //React.createElement("div",{value: option.id, id:option.name+"_button"},)
                  (<div id= {option.name}>{option.name}</div>)
              );
          })
          this.setState({menu : options});
          console.log("state", this.state.menu)
        })
    }
    render () {
      return (
        this.state.menu
        //this.state.menu
        //console.log(this)
        //this.parent.createElement("div",this.props,this.state.menu)
      );
    }
  }

  export default ChoiceButton