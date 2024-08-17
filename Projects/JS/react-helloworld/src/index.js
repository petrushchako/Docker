// Importing the necessary libraries and components
import React from 'react'; // React library is required to use JSX and create components
import ReactDOM from 'react-dom'; // ReactDOM is used to render React components into the DOM
import './App.css'; // Importing the CSS file to apply global styles to the application
import App from './App'; // Importing the main App component which contains the application logic

// Rendering the App component into the DOM
ReactDOM.render(
  // React.StrictMode is a wrapper component that helps identify potential problems
  <React.StrictMode>
    <App /> {/* App component is the root component of the application */}
  </React.StrictMode>,
  // Targeting the DOM element with id 'root' to render the React app
  document.getElementById('root')
);
