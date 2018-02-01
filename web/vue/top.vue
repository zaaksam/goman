<template>
    <div class="layout-ceiling">
        <div class="layout-logo">
            {{ title }}
        </div>
        <div class="layout-ceiling-main">
            <!-- <Avatar icon="person" size="small" /> -->
            <ButtonGroup>
                <Button icon="social-github" @click="onGithub"></Button>
                <Button @click="onLanguage">{{language}}</Button>
            </ButtonGroup>
        </div>
    </div>
</template>

<style>
.layout-ceiling {
  background: #464c5b;
  padding: 10px 0;
  overflow: hidden;
}

.layout-ceiling-main {
  float: right;
  margin-right: 15px;
}

.layout-ceiling-main a {
  color: #9ba7b5;
}

.layout-logo {
  width: 100px;
  height: 30px;
  line-height: 30px;
  color: aliceblue;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  left: 20px;
  text-align: center;
}
</style>

<script lang="ts">
import { Vue, Component, Watch } from "vue-property-decorator";
import iView from "iview";

@Component
export default class Top extends Vue {
  isCN = true;

  get title(): string {
    return C.appName + " " + C.appVersion;
  }

  get language(): string {
    return this.isCN ? "EN" : "中文";
  }

  // toggleClick() {
  //     if (this.app.spanLeft === 4) {
  //         this.app.spanLeft = 1
  //         this.app.spanRight = 23
  //     } else {
  //         this.app.spanLeft = 4
  //         this.app.spanRight = 20
  //     }
  // }

  onGithub() {
    window.open("https://github.com/zaaksam/goman");
  }

  onLanguage() {
    this.isCN = !this.isCN;

    let locale = this.isCN ? "zh-CN" : "en-US";
    this.$i18n.locale = locale;
  }
}
</script>
