// https://v3.nuxtjs.org/api/configuration/nuxt.config
import GlobalConfig from './global.json'

export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      Title: GlobalConfig.title,
      Description: GlobalConfig.description,
      Nav: GlobalConfig.nav,
      GetPageNumUrl: GlobalConfig.backend_url + "/ava-api/get-page-num",
      GetPostsUrl: GlobalConfig.backend_url + "/ava-api/get-posts",
      GetOnePostUrl: GlobalConfig.backend_url + "/ava-api/post/get-one",
      CheckUrl: GlobalConfig.backend_url + "/ava-api/check",
      LogOutUrl: GlobalConfig.backend_url + "/ava-api/logout",
      UploadImageUrl: GlobalConfig.backend_url + "/ava-api/image/upload",
      DeleteImageUrl: GlobalConfig.backend_url + "/ava-api/image/delete",
      UpdatePostUrl: GlobalConfig.backend_url + "/ava-api/post/update",
      DeletePostUrl: GlobalConfig.backend_url + "/ava-api/post/delete",
      GetImageBaseUrl: GlobalConfig.backend_url + "/ava-api/static/upload_images",
      GetPostsOfClassify: GlobalConfig.backend_url + "/ava-api/get-classify/get-posts",
      GetAllClassify: GlobalConfig.backend_url + "/ava-api/get-classify/get-all",
      GetPostsOfTimeline: GlobalConfig.backend_url + "/ava-api/get-timeline/get-posts",
      GetAllTimeline: GlobalConfig.backend_url + "/ava-api/get-timeline/get-all"
    }
  },

  compatibilityDate: "2024-07-29"
})
