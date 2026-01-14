const cardInner = document.getElementById('card-inner');
const inputNumber = document.getElementById('input-number');
const inputName = document.getElementById('input-name');
const inputExpiry = document.getElementById('input-expiry');
const inputCvv = document.getElementById('input-cvv');

// Displays
const displayNum = document.getElementById('display-number');
const displayName = document.getElementById('display-name');
const displayExpiry = document.getElementById('display-expiry');
const displayCvv = document.getElementById('display-cvv');
const displayBrand = document.getElementById('card-brand');

// Função para detectar bandeira
const detectBrand = (number) => {
    const cleaned = number.replace(/\D/g, '');
    
    // ELO: Diversos prefixos (simplificado)
    if (/^(4011|4312|4389|4514|4573|4576|5041|5066|5090|6277|6362|6363|6500|6504|6505|6516|6550)/.test(cleaned)) return 'ELO';
    
    // VISA: Começa com 4
    if (/^4/.test(cleaned)) return 'VISA';
    
    // MASTERCARD: 51-55 ou 2221-2720
    if (/^(5[1-5]|222[1-9]|22[3-9]|2[3-6]|27[01]|2720)/.test(cleaned)) return 'MASTERCARD';
    
    // AMEX: 34 ou 37
    if (/^3[47]/.test(cleaned)) return 'AMEX';
    
    
    // DINERS: 300-305, 36 ou 38
    if (/^3(?:0[0-5]|[68])/.test(cleaned)) return 'DINERS';

    // CIELO: Não é uma bandeira de emissão, mas sim uma adquirente. 
    // Geralmente cartões "Cielo" usam a base de outras bandeiras, 
    // mas se quiser detectar um range específico fictício para o projeto:
    if (/^60/.test(cleaned)) return 'CIELO';

    return 'BANK';
};

// Evento Número do Cartão
inputNumber.addEventListener('input', (e) => {
    let value = e.target.value.replace(/\D/g, '');
    value = value.replace(/(\d{4})(?=\d)/g, '$1 '); // Adiciona espaços
    e.target.value = value;
    
    displayNum.innerText = value || '#### #### #### ####';
   // displayBrand.innerText = detectBrand(value);

   const brand = detectBrand(value);
   updateCardStyle(brand);
});



// Evento Nome
inputName.addEventListener('input', (e) => {
    displayName.innerText = e.target.value.toUpperCase() || 'NOME DO TITULAR';
});

// Evento Validade
inputExpiry.addEventListener('input', (e) => {
    let value = e.target.value.replace(/\D/g, '');
    if (value.length > 2) {
        value = value.substring(0, 2) + '/' + value.substring(2, 4);
    }
    e.target.value = value;
    displayExpiry.innerText = value || 'MM/AA';
});

// Animação de Giro (Flip)
inputCvv.addEventListener('focus', () => {
    cardInner.classList.add('card-flipped');
});

inputCvv.addEventListener('blur', () => {
    cardInner.classList.remove('card-flipped');
});

inputCvv.addEventListener('input', (e) => {
    displayCvv.innerText = e.target.value || '***';
});

// Selecionar as faces do cartão
const cardFront = document.getElementById('card-front');
const cardBack = document.getElementById('card-back');

// Mapeamento de estilos por bandeira
const brandStyles = {
    'VISA': {
        gradient: ['from-blue-600', 'to-blue-900'],
        label: 'Visa'
    },
    'MASTERCARD': {
        // Usei este degradê linear para seguir seu padrão solicitado, 
        // mas você ainda pode usar a classe customizada se preferir
        gradient: ['from-red-600', 'to-yellow-500'],
        label: 'Mastercard'
    },
    'AMEX': {
        gradient: ['from-slate-100', 'to-zinc-400'], 
        label: 'American Express'
    },
    'ELO': {
        gradient: ['from-zinc-300', 'to-slate-950'],
        label: 'Elo'
    },
    'DINERS': {
        gradient: ['from-emerald-500', 'to-teal-700'],
        label: 'Diners Club'
    },
    'CIELO': {
        gradient: ['from-sky-400', 'to-blue-600'],
        label: 'Cielo'
    },
    'BANK': {
        gradient: ['from-gray-700', 'to-gray-900'],
        label: 'Bank'
    }
};

const updateCardStyle = (brand) => {
    const style = brandStyles[brand];
    
    // Remover classes de degradê antigas (limpeza)
    const allGradients = ['from-blue-600', 'to-blue-900', 'from-red-600', 'to-yellow-500', 'from-emerald-500', 'to-teal-700',
       'from-slate-100', 'to-zinc-400', 'from-zinc-300', 'to-slate-950','from-sky-400', 'to-blue-600','from-gray-700', 'to-gray-900'
    ];
    cardFront.classList.remove(...allGradients);
    cardBack.classList.remove(...allGradients);
    
    // Adicionar novas classes
    cardFront.classList.add('bg-gradient-to-br', style.gradient[0], style.gradient[1]);
    cardBack.classList.add('bg-gradient-to-br', style.gradient[0], style.gradient[1]);
    
    // Atualiza o texto da bandeira
    displayBrand.innerText = style.label;
};

const updateCardStyle2 = (brand) => {
    const style = brandStyles[brand] || brandStyles['BANK'];
    
    // Coleta todas as classes possíveis de todos os gradientes para limpar
    const allPossibleClasses = [];
    Object.values(brandStyles).forEach(b => {
        allPossibleClasses.push(...b.gradient);
    });

    // Remove classes antigas
    cardFront.classList.remove(...allPossibleClasses, 'bg-mastercard-custom'); // remove a custom se existir
    cardBack.classList.remove(...allPossibleClasses, 'bg-mastercard-custom');

    // Se for Mastercard e você quiser manter aquele efeito especial que criamos no CSS:
    if (brand === 'MASTERCARD') {
        cardFront.classList.add('bg-mastercard-custom');
        cardBack.classList.add('bg-mastercard-custom');
    } else {
        // Caso contrário, usa o degradê linear do Tailwind definido no objeto
        cardFront.classList.add('bg-gradient-to-br', style.gradient[0], style.gradient[1]);
        cardBack.classList.add('bg-gradient-to-br', style.gradient[0], style.gradient[1]);
    }

    displayBrand.innerText = style.label;
};