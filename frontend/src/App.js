
import './App.css';
import Header from "./components/Header/Header"
import SidePanel from './components/SidePanel/SidePanel';
import CreateMessage from './components/CreateMessage';

function App() {

    return (
        <div>
          <Header/>
          <SidePanel  />
          <CreateMessage />

        </div>
    );
}

export default App;
