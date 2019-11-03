import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import JB from './App';
//import ChoiceButton from './GameButton';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
  ReactDOM.render(<JB />, div);
  ReactDOM.unmountComponentAtNode(div);
});
