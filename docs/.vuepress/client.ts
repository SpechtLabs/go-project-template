import { defineClientConfig } from "vuepress/client";

export default defineClientConfig({
  enhance() {
    // The shared components (Terminal, FileTree, ListCompare, Contributors,
    // Releases and their VP* home page sections) come from
    // @spechtlabs/docs-kit, which registers them itself (see config.ts).
  },
});
