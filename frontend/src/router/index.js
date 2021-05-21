import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from "../views/index/index"
import Posts from "../views/index/posts"
import OnePost from "../views/index/one-post"
import Admin from "../views/admin/admin"
import Modify from "../views/admin/modify"
import Manage from "../views/admin/manage"
import Check from "../views/admin/check"
import PageNotFound from "../views/index/page-not-found"
import Classify from "../views/index/classify"
import TimeLine from "../views/index/timeline"

Vue.use(VueRouter)

export default new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes: [
		{
			path: '/',
			component: Index,
			children: [
				{
					path: 'post/:id',
					name: 'Post',
					component: OnePost
				},
				{
					path: '',
					name: 'Posts',
					component: Posts
				},
				{
					path: 'page-not-found',
					name: 'PageNotFound',
					component: PageNotFound
				},
				{
					path: 'classify',
					name: 'Classify',
					component: Classify
				},
				{
					path: 'timeline',
					name: 'TimeLine',
					component: TimeLine
				}
			]
		},
		{
			path: '/admin',
			name: 'Admin',
			component: Admin,
			children: [
				{
					path: 'manage',
					name: 'Manage',
					component: Manage
				},
				{
					path: 'modify/:id',
					name: 'New',
					component: Modify
				},
				{
					path: 'check',
					name: 'Check',
					component: Check
				},
			]
		},
		{
			path: "*",
			redirect: "/page-not-found"
		}
	]
})
