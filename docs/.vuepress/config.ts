import { docsKitPlugin } from "@spechtlabs/docs-kit";
import { viteBundler } from "@vuepress/bundler-vite";
import { defineUserConfig } from "vuepress";
import { plumeTheme } from "vuepress-theme-plume";

export default defineUserConfig({
  base: "/",
  lang: "en-US",
  title: "PROJECT_NAME",
  description: "PROJECT_NAME does one thing, and does it well",

  head: [
    ["meta", { name: "description", content: "PROJECT_NAME does one thing, and does it well" }],
    ["link", { rel: "icon", type: "image/png", href: "/images/specht.png" }],
  ],

  bundler: viteBundler(),
  shouldPrefetch: false,

  plugins: [
    // The components and Markdown containers shared by the SpechtLabs docs
    // sites (::: terminal, ::: cast, FileTree, Contributors, Releases, ...).
    // https://github.com/SpechtLabs/docs-kit
    docsKitPlugin({
      github: {
        // Every repository listed here is fetched from the GitHub API at build
        // time, and in CI a failed fetch fails the build. A repository that
        // doesn't exist yet (the template's placeholder name) or is still
        // private would break it, so the list starts empty. Once the
        // repository is public, add "SpechtLabs/PROJECT_NAME" here to use the
        // VPContributors and VPReleases home page sections.
        repos: [],
      },
    }),
  ],

  theme: plumeTheme({
    docsRepo: "https://github.com/SpechtLabs/PROJECT_NAME",
    docsDir: "docs",
    docsBranch: "main",

    editLink: true,
    lastUpdated: false,
    contributors: false,

    blog: false,

    cache: "filesystem",
    search: { provider: "local" },

    sidebar: {
      // Getting Started: the tutorial, which teaches by doing.
      "/getting-started/": [
        {
          text: "Getting Started",
          icon: "mdi:rocket-launch",
          prefix: "/getting-started/",
          items: [{ text: "Quick start", link: "quick-start", icon: "mdi:flash", badge: "5 min" }],
        },
      ],

      // How-to Guides: one task each.
      "/guides/": [
        {
          text: "How-to Guides",
          icon: "mdi:compass",
          prefix: "/guides/",
          items: [{ text: "Install PROJECT_NAME", link: "install", icon: "mdi:download" }],
        },
      ],

      // Understanding: background and design.
      "/understanding/": [
        {
          text: "Understanding",
          icon: "mdi:lightbulb",
          prefix: "/understanding/",
          items: [{ text: "Overview", link: "overview", icon: "mdi:eye" }],
        },
      ],

      // Reference: lookup only.
      "/reference/": [
        {
          text: "Reference",
          icon: "mdi:book",
          prefix: "/reference/",
          items: [{ text: "Command line", link: "cli", icon: "mdi:terminal" }],
        },
      ],
    },

    markdown: {
      collapse: true,
      mermaid: true,
    },

    watermark: false,
  }),
});
