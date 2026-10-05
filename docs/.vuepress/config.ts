import { viteBundler } from "@vuepress/bundler-vite";
import { registerComponentsPlugin } from "@vuepress/plugin-register-components";
import { path } from "@vuepress/utils";
import container from "markdown-it-container";
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

  // ::: terminal blocks render as a terminal window (components/Terminal.vue).
  extendsMarkdown: (md) => {
    md.use(container, "terminal", {
      validate: (params: string) => {
        const info = params.trim();
        return /^terminal(?:\s+.*)?$/.test(info);
      },
      render: (tokens: any[], idx: number) => {
        const token = tokens[idx];
        if (token.nesting === 1) {
          const info = token.info.trim();
          const rest = info.replace(/^terminal\s*/, "");
          const attrs: Record<string, string> = {};
          const attrRegex = /(\w+)=((?:\"[^\"]*\")|(?:'[^']*')|(?:[^\s]+))/g;
          let consumed = "";
          let m: RegExpExecArray | null;
          while ((m = attrRegex.exec(rest)) !== null) {
            const key = m[1];
            let val = m[2];
            if ((val.startsWith('"') && val.endsWith('"')) || (val.startsWith("'") && val.endsWith("'"))) {
              val = val.slice(1, -1);
            }
            attrs[key] = val;
            consumed += m[0] + " ";
          }
          const positional = rest.replace(consumed, "").trim();
          const titleRaw = attrs.title ?? positional ?? "";
          const title = titleRaw ? md.utils.escapeHtml(titleRaw) : "";
          const titleAttr = title ? ` title=\"${title}\"` : "";
          return `\n<Terminal${titleAttr}>\n`;
        }
        return `\n</Terminal>\n`;
      },
    });
  },

  plugins: [
    registerComponentsPlugin({
      componentsDir: path.resolve(__dirname, "./components"),
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
