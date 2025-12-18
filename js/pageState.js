// Page State Management
export const pageState = {
    currentPage: 'home',
    isAnimating: false,
    animationDirection: 'forward',
    pageHistory: ['home'],
    pages: [
        { id: 'home', title: 'Home', order: 0 },
        { id: 'about', title: 'About', order: 1 },
        { id: 'skills', title: 'Skills', order: 2 },
        { id: 'achievements', title: 'Achievements', order: 3 },
        { id: 'projects', title: 'Projects', order: 4 },
        { id: 'contact', title: 'Contact', order: 5 }
    ]
};

export function updatePageState(newPageId) {
    if (pageState.currentPage !== newPageId) {
        pageState.pageHistory.push(newPageId);
        if (pageState.pageHistory.length > 10) {
            pageState.pageHistory.shift();
        }
    }
    pageState.currentPage = newPageId;
}

export function getCurrentPageOrder() {
    return getPageOrder(pageState.currentPage);
}

export function getPageOrder(pageId) {
    const page = pageState.pages.find(p => p.id === pageId);
    return page ? page.order : 0;
}