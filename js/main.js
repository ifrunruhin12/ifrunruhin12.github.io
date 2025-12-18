// Main JavaScript Entry Point - Newspaper Portfolio
import { audioSystem } from './audioSystem.js';
import { initializeNavigation, initializeMobileMenu, initializeKeyboardNavigation } from './navigation.js';

// Initialize the application
document.addEventListener('DOMContentLoaded', function() {
    // Initialize audio system first
    audioSystem.init();
    
    // Initialize navigation systems
    initializeNavigation();
    initializeMobileMenu();
    initializeKeyboardNavigation();
    
    // Add page turn indicator
    addPageTurnIndicator();
    
    // Optimize performance
    optimizePageFlipPerformance();
});

function addPageTurnIndicator() {
    const indicator = document.createElement('div');
    indicator.className = 'page-turn-indicator';
    indicator.innerHTML = '📖';
    indicator.title = 'Use arrow keys or click navigation to flip pages';
    document.body.appendChild(indicator);
    
    // Show indicator briefly on load
    setTimeout(() => {
        indicator.classList.add('visible');
        setTimeout(() => {
            indicator.classList.remove('visible');
        }, 3000);
    }, 1000);
    
    // Show on hover over navigation
    const navLinks = document.querySelectorAll('.nav-link');
    navLinks.forEach(link => {
        link.addEventListener('mouseenter', () => {
            indicator.classList.add('visible');
        });
        link.addEventListener('mouseleave', () => {
            indicator.classList.remove('visible');
        });
    });
}

function optimizePageFlipPerformance() {
    // Preload all pages for smoother transitions
    const pages = document.querySelectorAll('.newspaper-page');
    pages.forEach(page => {
        // Force hardware acceleration
        page.style.transform = 'translateZ(0)';
        page.style.willChange = 'transform';
        
        // Preload images
        const images = page.querySelectorAll('img');
        images.forEach(img => {
            if (img.src) {
                const preloadImg = new Image();
                preloadImg.src = img.src;
            }
        });
    });
    
    // Enable audio on first user interaction
    document.addEventListener('click', function enableAudio() {
        if (audioSystem.audio) {
            audioSystem.audio.play().then(() => {
                audioSystem.audio.pause();
                audioSystem.audio.currentTime = 0;
            }).catch(() => {
                // Audio will be enabled on next interaction
            });
        }
        document.removeEventListener('click', enableAudio);
    }, { once: true });
}