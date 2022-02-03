export interface Article {
    title: string;
    tags: string[];
    date: Date;
    src: string;
    url: string;
}

export function GetArticles(from?: string): Promise<Article[]> {
    const target = `http://localhost:8080/news/${from ? from : "" }`

    return fetch(target)
        .then((res) => res.json())
        .then((res) => res.map((article: any) => formatArticle(article)));
}

function formatArticle(article: any): Article {
    return { title: article.title, tags: article.tags, date: article.date, src: article.src, url: article.url };
}
