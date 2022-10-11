import {useFetch} from "nuxt/app";

function setIfChecked(t) {
  useFetch(CheckUrl, {
    method: "post"
  }).then(res => {
    if (t.user !== undefined) t.user = res.data.user
      t.isChecked = res.data.isChecked
    })
}

function logout(t) {
  useFetch(LogOutUrl, {
    method: "get"
  }).then(res => {
    t.$router.go(0)
  })
}

function toCheck(t) {
  t.$router.push('/admin/check')
}

export {
  logout,
  setIfChecked,
  toCheck
}