
/**
 * Checks if an object is empty
 * @param obj - Object ot check
 * @returns {boolean} - True if object is empty
 */
export function isObjectEmpty(obj) {
    for(let i in obj) return false;
    return true;
}

/**
 * Checks if object is a valid date
 * @param d - object to check
 * @returns {boolean} - True if object is a valid date
 */
export function isValidDate(d) {
    return d instanceof Date && !isNaN(d);
}

/**
 * Check if a string canbe coverted to a valid date
 * @param stringDate - String possibly containing a date
 * @returns {boolean} - True if string contains a valid date
 */
export function isValidStringDate(stringDate) {
    if (stringDate === "") {
        return false
    }
    let d = new Date(stringDate)
    return isValidDate(d)
}

/**
 * Take an object and convert it to name, value KV pairs
 * @param obj - Object ot convert
 * @returns {{string, string}[]} - KV pairs
 */
export function convertExtraFields(obj) {
    let vals = []
    let key;
    if (!isObjectEmpty(obj)) {
        for (let key1 in obj) {
            key = key1
            vals.push({'name': key, 'value': obj[key]})
        }
        return vals
    }
}

/**
 * Since some fields come as date time and we need to display them as dates,
 * we need to update them manually
 * @param dateToUpdate - String, name of the date field to update
 * @param newDateString - String that contains the date to be used for update
 * @param objectToUpdate - Object to be updated
 */
export function updateValidDate(dateToUpdate, newDateString, objectToUpdate){
    let foundVD = false
    if (objectToUpdate) {
        if (objectToUpdate.validityDates) {
            let parts = newDateString.split('-');
            let VD = new Date(parts[0], parts[1] - 1, parts[2]);
            if (isValidDate(VD)){
                objectToUpdate.validityDates[dateToUpdate] = VD
            }
            foundVD = true
        }
    }
    if (!foundVD) {
        alert('Unable to populate validity date')
    }
}