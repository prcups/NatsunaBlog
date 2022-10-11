// https://v3.nuxtjs.org/api/configuration/nuxt.config
import GlobalConfig from './global.json'

export default defineNuxtConfig({
    runtimeConfig: {
        Title: GlobalConfig.title,
        Description: GlobalConfig.description,
        Nav: GlobalConfig.nav,
        GetPageNumUrl: GlobalConfig.base_url + "/ava-api/get-page-num",
        GetPostsUrl: GlobalConfig.base_url + "/ava-api/get-posts",
        GetOnePostUrl: GlobalConfig.base_url + "/ava-api/post/get-one",
        CheckUrl: GlobalConfig.base_url + "/ava-api/check",
        LogOutUrl: GlobalConfig.base_url + "/ava-api/logout",
        UploadImageUrl: GlobalConfig.base_url + "/ava-api/image/upload",
        DeleteImageUrl: GlobalConfig.base_url + "/ava-api/image/delete",
        UpdatePostUrl: GlobalConfig.base_url + "/ava-api/post/update",
        DeletePostUrl: GlobalConfig.base_url + "/ava-api/post/delete",
        GetImageBaseUrl: GlobalConfig.base_url + "/static/upload_images",
        GetPostsOfClassify: GlobalConfig.base_url + "/ava-api/get-classify/get-posts",
        GetAllClassify: GlobalConfig.base_url + "/ava-api/get-classify/get-all",
        GetPostsOfTimeline: GlobalConfig.base_url + "/ava-api/get-timeline/get-posts",
        GetAllTimeline: GlobalConfig.base_url + "/ava-api/get-timeline/get-all"
    }
})
