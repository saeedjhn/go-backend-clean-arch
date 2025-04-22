# 🌐 Multi-language (i18n) Support in Web Projects

In a project, **multi-language** can mean:

- Supporting multiple spoken/written languages (**internationalization** or **i18n**).
- Using multiple programming languages (not the focus here).

Since this is in the context of frontend vs backend, we're talking about **i18n**—supporting multiple human languages in
your app.

---

## 🧠 Where to Handle Multi-language Support?

### ✅ Frontend: **Preferred place for i18n**

#### Why?

- User-facing text lives here.
- Easier to switch languages dynamically (e.g., dropdown to switch to Spanish).
- Works great with translation files (e.g., JSON).

#### 🔧 Common Frontend Libraries:

- **React**: `react-i18next`, `next-translate`
- **Vue**: `vue-i18n`
- **Angular**: built-in i18n support

#### ✅ Pros:

- No need to reload the page for switching languages.
- Faster user experience—translations are already loaded.
- You can lazy-load translations per language/page.

---

### 🛠 Backend: **Optional, only if needed**

#### Use when:

- Sending **language-specific content** (e.g., emails, PDFs).
- Providing **default translations** in APIs.
- Centralized **fallback logic** is required (e.g., for logs, server-rendered pages).

#### ✅ Pros:

- Consistent logic across services.
- Good for SSR or mobile apps that consume your API.

#### 🧰 Common Backend Libraries:

- **Node.js**: `i18next`, `node-polyglot`
- **Go**: `nicksnyder/go-i18n`
- **Python**: `gettext`, `Babel`
- **Java**: `ResourceBundle`

---

## ✅ Best Practice

> Let the **frontend** handle all UI translations.  
> Let the **backend** support localization only when returning specific content (emails, PDFs, etc.).

### 📩 Tips:

- Use `Accept-Language` HTTP header to inform backend of the preferred language.
- Use consistent locale formats (`en-US`, `fa-IR`, etc.).

---

## 📦 Want Examples?

Need a full code example or setup for your stack (React, Go, NestJS, etc.)?  
Let me know and I’ll generate a boilerplate for you! 😄
