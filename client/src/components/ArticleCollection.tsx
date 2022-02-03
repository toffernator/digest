import { useEffect, useState } from "react";
import Item from "./Item";
import ItemCollection from "./ItemCollection";
import { Article, GetArticles } from "../digest-client/client"

export default function ArticleCollection() {
    const emptyArticles: Article[] = [];
    const [articles, setArticles] = useState(emptyArticles);
    useEffect(() => {
        GetArticles().then((articles) => setArticles(articles))
    }, [])

    const items = articles.map((article) => <Item title = {article.title} tags = {article.tags} src = {article.src} url = {article.url}/>)
    return <ItemCollection items={items} />
}
