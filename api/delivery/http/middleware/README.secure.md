## Content Security Policy (CSP)

### Overview

A Content Security Policy (CSP) is a security feature that helps protect web applications from various attacks like
Cross-Site Scripting (XSS), clickjacking, and data injection. By specifying allowed content sources, CSPs minimize the
risk of executing malicious content on your site.

### CSP Guidance

#### AppKit

The following is a partial CSP that covers WalletConnect's libraries and services for AppKit. Note that you may need to
define additional sources based on your application's requirements.

```
default-src 'self';
script-src 'self';
style-src https://fonts.googleapis.com;
img-src * 'self' data: blob: https://walletconnect.org https://walletconnect.com https://secure.walletconnect.com https://secure.walletconnect.org https://tokens-data.1inch.io https://tokens.1inch.io https://ipfs.io;
font-src 'self' https://fonts.gstatic.com;
connect-src 'self' https://rpc.walletconnect.com https://rpc.walletconnect.org https://relay.walletconnect.com https://relay.walletconnect.org wss://relay.walletconnect.com wss://relay.walletconnect.org https://pulse.walletconnect.com https://pulse.walletconnect.org https://api.web3modal.com https://api.web3modal.org https://keys.walletconnect.com https://keys.walletconnect.org https://notify.walletconnect.com https://notify.walletconnect.org https://echo.walletconnect.com https://echo.walletconnect.org https://push.walletconnect.com https://push.walletconnect.org wss://www.walletlink.org;
frame-src 'self' https://verify.walletconnect.com https://verify.walletconnect.org https://secure.walletconnect.com https://secure.walletconnect.org;
```

---

### go-echo framework

#### Secure

Secure middleware provides protection against cross-site scripting (XSS) attack, content type sniffing, clickjacking,
insecure connection and other code injection attacks.

1. XSSProtection
    - Explanation: This header tells browsers to enable XSS (Cross-Site Scripting) protection. It attempts to block
      harmful scripts that can be injected into a page.
    - Recommended Value: It should generally be set to "1; mode=block" to ensure that any potential XSS attack is
      blocked and the page is not rendered if an attack is detected.
        - 1: Activates XSS protection.
        - mode=block: If an XSS attack is detected, the browser should block the page rendering.


2. ContentTypeNosniff
    - Explanation: This header prevents browsers from sniffing the content type and ensures that the browser adheres to
      the specified Content-Type header. It helps prevent attacks where the browser might misinterpret a file's content
      type.
    - Recommended Value: It should be set to "true", instructing browsers to not attempt to sniff the content type.


3. XFrameOptions
    - Explanation: This header tells browsers to avoid displaying the page in a frame. This is crucial to prevent
      Clickjacking, where a malicious site might embed your page inside a frame and trick users into clicking on hidden
      buttons or links.
    - Recommended Value: The best value is "DENY", which prevents the page from being shown in any frame. If you only
      want the page to be allowed in frames from the same origin, you can use "SAMEORIGIN".


4. HSTSMaxAge (HTTP Strict Transport Security)
    - Explanation: The HSTS header forces browsers to only use HTTPS for subsequent requests to your site for a
      specified period. This helps prevent SSL Stripping attacks, where an attacker downgrades the connection from HTTPS
      to HTTP.
    - Recommended Value: It should be set to a reasonable duration, such as one year (31,536,000 seconds). During
      development, you can use a smaller value like one hour (3600 seconds).


5. HSTSExcludeSubdomains (HTTP Strict Transport Security (HSTS))
   The HSTSExcludeSubdomains option is related to the HTTP Strict Transport Security (HSTS) header, which enforces the
   use of HTTPS (secure connections) for your site. Here's an explanation of this option:

    - Explanation:
        - HSTS (HTTP Strict Transport Security) is a security feature that tells browsers to only access a site using
          HTTPS for a specified period of time (defined by the HSTSMaxAge header). It helps protect against attacks like
          SSL
          Stripping, where an attacker might try to downgrade a secure HTTPS connection to an insecure HTTP one.

        - The option HSTSExcludeSubdomains controls whether the HSTS policy should apply to all subdomains of the site.
          By default, HSTS applies to the domain and all of its subdomains. However, if this option is enabled (true),
          the HSTS policy will not apply to subdomains of the current domain.

    - Behavior:
        - When HSTSExcludeSubdomains is true, it ensures that subdomains are excluded from the HSTS policy, even if the
          main
          domain has HSTS enabled. This means subdomains won't be forced to use HTTPS as per the policy.

        - The option has no effect unless the HSTSMaxAge is set to a non-zero value. This is because HSTS itself is only
          relevant when you set an expiration time (HSTSMaxAge) for the policy.

    - Default Value:
      The default value for HSTSExcludeSubdomains is false, meaning HSTS will be applied to the main domain and all its
      subdomains unless specified otherwise.

    - Summary:
        - HSTSExcludeSubdomains: When set to true, it excludes subdomains from the HSTS policy.
        - Effect: It only works if HSTSMaxAge is set to a non-zero value, which is required for HSTS to be enforced.

6. ContentSecurityPolicy (CSP)
    - Explanation: Content Security Policy (CSP) allows you to control which resources can be loaded on your page (such
      as scripts, images, styles). This helps protect against attacks like Cross-Site Scripting (XSS) and Data
      Injection. By restricting the sources from which your site can load content, you reduce the chances of malicious
      content being executed.
    - Recommended Value: You should start with a restrictive policy like "default-src 'self'", which ensures that only
      resources from the same origin can be loaded. You can further specify individual content types (e.g., scripts,
      images) and where they are allowed to come from.

#### Recommended example

```go
e := echo.New()
e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
XSSProtection:         "1; mode=block",
ContentTypeNosniff:    "true",
XFrameOptions:         "DENY",
HSTSMaxAge:            31536000, // 1 year
HSTSExcludeSubdomains: true,        // Excludes subdomains from the HSTS policy
ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self';",
}))

```