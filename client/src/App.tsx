import ArticleCollection from "./components/ArticleCollection";
import "./index.css";

function App() {
    return (
        <div className="container mx-auto">
            <h1 className="text-9xl text-center">News</h1>
            <ArticleCollection />
        </div>
    );
}

export default App;
