// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    publicRuntimeConfig: {
        GetPageNumUrl: process.env.BASE_URL + "/ava-api/get-page-num",
        GetPostsUrl: process.env.BASE_URL + "/ava-api/get-posts",
        GetOnePostUrl: process.env.BASE_URL + "/ava-api/post/get-one",
        CheckUrl: process.env.BASE_URL + "/ava-api/check",
        LogOutUrl: process.env.BASE_URL + "/ava-api/logout",
        UploadImageUrl: process.env.BASE_URL + "/ava-api/image/upload",
        DeleteImageUrl: process.env.BASE_URL + "/ava-api/image/delete",
        UpdatePostUrl: process.env.BASE_URL + "/ava-api/post/update",
        DeletePostUrl: process.env.BASE_URL + "/ava-api/post/delete",
        GetImageBaseUrl: process.env.BASE_URL + "/static/upload_images",
        GetPostsOfClassify: process.env.BASE_URL + "/ava-api/get-classify/get-posts",
        GetAllClassify: process.env.BASE_URL + "/ava-api/get-classify/get-all",
        GetPostsOfTimeline: process.env.BASE_URL + "/ava-api/get-timeline/get-posts",
        GetAllTimeline: process.env.BASE_URL + "/ava-api/get-timeline/get-all"
    }
})
