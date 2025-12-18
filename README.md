# Newspaper Portfolio

A unique portfolio website designed to mimic the appearance and navigation of a traditional newspaper, featuring smooth page-flipping animations and authentic newspaper aesthetics.

## 🏗️ Architecture

This project uses a **modular architecture** for better maintainability and organization:

### CSS Structure (`css/`)
- **`main.css`** - Entry point that imports all CSS modules
- **`variables.css`** - CSS custom properties and responsive breakpoints
- **`base.css`** - Reset styles, typography, and base elements
- **`layout.css`** - Grid system and container layouts
- **`navigation.css`** - Navigation menu and mobile menu styles
- **`pages.css`** - Page system and basic page animations
- **`components.css`** - Reusable newspaper components (articles, masthead, etc.)
- **`animations.css`** - Advanced page flip animations and effects

### JavaScript Structure (`js/`)
- **`main.js`** - Entry point and application initialization
- **`pageState.js`** - Page state management and utilities
- **`audioSystem.js`** - Audio system for page flip sounds
- **`navigation.js`** - Navigation logic and event handlers
- **`animations.js`** - Page flip animation system

## � Featurces

- **Authentic Newspaper Design**: Black and white color scheme, serif typography, column layouts
- **Smooth Page Flip Animations**: 3D CSS transforms with realistic physics
- **Responsive Design**: Adapts to desktop, tablet, and mobile devices
- **Audio Feedback**: Page turn sound effects
- **Keyboard Navigation**: Arrow keys for page navigation
- **Mobile-Friendly**: Touch-friendly interactions and collapsible menu
- **Performance Optimized**: Hardware acceleration and efficient animations

## 📱 Responsive Breakpoints

- **Desktop**: 1000px width, 1200px height
- **Tablet** (≤768px): 100% width, 900px height
- **Mobile** (≤480px): 100% width, 700px height

## 🎨 Design System

### Typography
- **Headlines**: Times New Roman (serif)
- **Body Text**: Arial (sans-serif)
- **Monospace**: Courier New

### Color Palette
- **Primary**: Black (#000000)
- **Background**: White (#ffffff)
- **Secondary**: Dark Gray (#333333)
- **Accent**: Light Gray (#666666)

## 🛠️ Development

### File Structure
```
├── index.html              # Main HTML file
├── css/                    # Modular CSS files
│   ├── main.css           # CSS entry point
│   ├── variables.css      # CSS custom properties
│   ├── base.css          # Base styles and typography
│   ├── layout.css        # Layout system
│   ├── navigation.css    # Navigation styles
│   ├── pages.css         # Page system
│   ├── components.css    # Newspaper components
│   └── animations.css    # Page flip animations
├── js/                     # Modular JavaScript files
│   ├── main.js           # JavaScript entry point
│   ├── pageState.js      # State management
│   ├── audioSystem.js    # Audio system
│   ├── navigation.js     # Navigation logic
│   └── animations.js     # Animation system
└── assets/                # Static assets
    ├── profile.jpg       # Profile image
    ├── cv.pdf           # CV document
    └── page-turn.mp3    # Page flip sound
```

### Browser Support
- Modern browsers with ES6 module support
- CSS Grid and Flexbox support
- 3D CSS transforms (with fallbacks)

## 📄 Pages

1. **Home** - Welcome page with masthead and introduction
2. **About** - Developer profile and background
3. **Skills** - Technical skills in classified format
4. **Achievements** - Accomplishments as news articles
5. **Projects** - Project portfolio as feature articles
6. **Contact** - Contact information in classified format

## 🔧 Customization

### Adding New Pages
1. Add page data to `pageState.pages` in `js/pageState.js`
2. Create HTML section in `index.html`
3. Add navigation link in the nav menu

### Modifying Styles
- Edit the appropriate CSS module in the `css/` directory
- Variables can be changed in `css/variables.css`
- Component styles are in `css/components.css`

### Performance Notes
- Uses CSS `will-change` for animation optimization
- Hardware acceleration with `translateZ(0)`
- Image preloading for smooth transitions
- Efficient event handling with delegation

## 📝 License

This project is open source and available under the MIT License.