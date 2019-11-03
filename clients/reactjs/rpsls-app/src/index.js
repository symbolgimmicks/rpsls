import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import * as gameButtons from "./GameButton";
//import ChoiceButton as gameButtons from "./GameButton";

ReactDOM.render(<App />, document.getElementById('root'));
//ReactDOM.render(<gameButtons.ChoiceButton />, document.getElementById('buttonTest'));
ReactDOM.render(<gameButtons.SelectADecision />, document.getElementById('menuTest'));
// ReactDOM.render(<gameButtons.rockButton />, document.getElementById('menuTest'));
// ReactDOM.render(<gameButtons.paperButton />, document.getElementById('menuTest'));
// ReactDOM.render(<gameButtons.scissorsButton />, document.getElementById('menuTest'));
// ReactDOM.render(<gameButtons.lizardButton />, document.getElementById('menuTest'));
// ReactDOM.render(<gameButtons.spockButton />, document.getElementById('menuTest'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
