package news

const (
	DR = "dr"
)

const (
	DR_URL = "https://www.dr.dk/tjenester/steffi/graphql?query=fragment%20ArticleTeaser%20on%20Article%20%7B%20type%3A%20__typename%20urn%20url%20title%20format%20publications%20%7B%20...on%20ArticlePublication%20%7B%20breaking%20live%20serviceChannel%20%7B%20urn%20%7D%20%7D%20%7D%20summary%20startDate%20teaserImage%20%7B%20default%20%7B%20url%20%7D%20%7D%20teaserVideo%20%7B%20type%3A%20__typename%20resource%20%7B%20type%3A%20__typename%20...%20on%20Channel%20%7B%20slug%20%7D%20...%20on%20ProgramCard%20%7B%20urn%20%7D%20...%20on%20ProgramCardBundle%20%7B%20items(limit%3A%201)%20%7B%20__typename%20urn%20%7D%20%7D%20%7D%20%7D%20head%20%7B%20type%3A%20__typename%20...%20on%20MediaComponent%20%7B%20resource%20%7B%20type%3A%20__typename%20...%20on%20ProgramCard%20%7B%20primaryAsset%20%7B%20durationInMilliseconds%20%7D%20%7D%20...%20on%20ProgramCardBundle%20%7B%20items(limit%3A%201)%20%7B%20primaryAsset%20%7B%20durationInMilliseconds%20%7D%20%7D%20%7D%20%7D%20%7D%20...%20on%20ImageCollectionComponent%20%7B%20images%20%7B%20default%3A%20image(key%3A%20%22default%22)%20%7B%20type%3A%20__typename%20%7D%20%7D%20%7D%20%7D%20contributions(limit%3A%201)%20%7B%20agent%20%7B%20...%20on%20Person%20%7B%20name%20%7D%20%7D%20%7D%20site%20%7B%20url%20urn%20title%20presentation%20%7B%20colors%20%7D%20%7D%20%7D%20query%20NewsOverview%20%7B%20frontPage(id%3A%20%225d023c534ff3c845599d0953%22)%20%7B%20topStories%20%7B%20url%20title%20live%20channel%20%7B%20slug%20%7D%20image%20%7B%20url%20width%20height%20%7D%20article%20%7B%20...ArticleTeaser%20%7D%20%7D%20%7D%20%7D"
)
