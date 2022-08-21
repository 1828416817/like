<template>
  <div class="article-reward " style="caret-color: transparent">
      <a :class="isLike" @click="like">
        <v-icon size="14" color="#fff">mdi-thumb-up</v-icon> 点赞
        <span v-show="article.likeCount > 0" >{{
            article.likeCount
          }}</span>
      </a>
    </div>
</template>

<script>
import axios from "axios"
export default {
  data() {
    return {
      article:{
        likeCount: 12
      },
      isLike: 'like-btn-active',
    }
  },
  methods: {
    like(){
      axios.post("http://localhost:9000/like/1").then(res => {
        if (res){
          alert("success")
        }
      })
      if (this.isLike==='like-btn-active'){
        this.isLike='like-btn'
        this.article.likeCount=this.article.likeCount-1;
      }else {
        this.isLike = 'like-btn-active'
        this.article.likeCount=this.article.likeCount+1;
      }
    }
  }
}
</script>

<style  scoped type="text/css">
.like-btn-active {
  display: inline-block;
  width: 100px;
  background: #ec7259;
  color: #fff !important;
  text-align: center;
  line-height: 36px;
  font-size: 0.875rem;
}
.like-btn {
  display: inline-block;
  width: 100px;
  background: #969696;
  color: #fff !important;
  text-align: center;
  line-height: 36px;
  font-size: 0.875rem;
}
.article-reward {
  margin-top: 5rem;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>

