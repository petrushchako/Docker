// Importing necessary libraries and components
import React from 'react'; // React is a library for building user interfaces
import './App.css'; // Importing the CSS file to style the components

// Defining a functional component called App
function App() {
  // The return statement defines the JSX (JavaScript XML) to be rendered
  return (
    <div className="App">
      {/* The <header> element represents the header of the application */}
      <header className="App-header">
        {/* Displaying the Hello World message */}
        <h1>Hello World!</h1>
        {/* A paragraph element with some additional text */}
        <p>
          Welcome to my first React app!
        </p>
      </header>
    </div>
  );
}

// Exporting the App component as the default export
// This allows the App component to be imported and used in other files
export default App;
