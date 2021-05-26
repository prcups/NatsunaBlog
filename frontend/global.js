var baseUrl = "http://127.0.0.1:8199"
var title = "Yuno"
var description = "再年轻一次吧"
var tag = "博客,龙芯,C++,Go,Rust"
var onNav = [
    {
        "title": "关于",
        "url": "https://mirror.lzu.edu.cn"
    }
]

export default {
    Title: title,
    Description: description,
    Tag: tag,
    OnNav: onNav,

    GetPageNumUrl: baseUrl + "/ava-api/get-page-num",
    GetPostsUrl: baseUrl + "/ava-api/get-posts",
    GetOnePostUrl: baseUrl + "/ava-api/post/get-one",
    CheckUrl: baseUrl + "/ava-api/check",
    LogOutUrl: baseUrl + "/ava-api/logout",
    UploadImageUrl: baseUrl + "/ava-api/image/upload",
    DeleteImageUrl: baseUrl + "/ava-api/image/delete",
    UpdatePostUrl: baseUrl + "/ava-api/post/update",
    DeletePostUrl: baseUrl + "/ava-api/post/delete",
    GetImageBaseUrl: baseUrl + "/static/upload_images",
    GetPostsOfClassify: baseUrl + "/ava-api/get-classify/get-posts",
    GetAllClassify: baseUrl + "/ava-api/get-classify/get-all",
    GetPostsOfTimeline: baseUrl + "/ava-api/get-timeline/get-posts",
    GetAllTimeline: baseUrl + "/ava-api/get-timeline/get-all"
}