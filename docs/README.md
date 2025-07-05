---
pageLayout: home
externalLinkIcon: false

config:
  - type: doc-hero
    hero:
      name: {{REPO_NAME}}
      text: {{REPO_TAGLINE}}
      tagline: {{REPO_DESCRIPTION}}
      image: /logo.png
      actions:
        - text: Get Started →
          link: /guide/overview
          theme: brand
          icon: simple-icons:bookstack
        - text: GitHub Releases →
          link: https://github.com/SpechtLabs/{{REPO_NAME}}/releases
          theme: alt
          icon: simple-icons:github

  - type: features
    title: Why {{REPO_NAME}}?
    description: {{REPO_DESCRIPTION}}
    features:
      - title: 1
        icon: mdi:shield-lock-outline
        details: No public endpoints. Access is gated by your Tailscale ACLs, identity, and devices — and nothing else.

      - title: 2
        icon: mdi:timer-sand
        details: All kubeconfigs are short-lived and auto-expiring. You get access when you need it, and not a second longer.

      - title: 3
        icon: mdi:kubernetes
        details: Grants are mapped to native ClusterRoles, and credentials are provisioned as real Kubernetes ServiceAccounts.

      - title: 4
        icon: mdi:account-key-outline
        details: Use a CRD to define how Tailscale users, groups, or tags map to Kubernetes roles — GitOps ready.

  - type: VPReleasesCustom
    repo: SpechtLabs/{{REPO_NAME}}

  - type: VPContributorsCustom
    repo: SpechtLabs/{{REPO_NAME}}
---
